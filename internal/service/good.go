package service

import (
	"context"
	"io"
	"path"
	"strconv"
	"strings"

	pb "valuation/api/valuation/v1"
	"valuation/internal/biz"
	"valuation/internal/data"
	"valuation/pkg/convertx"
	"valuation/pkg/errorx"
	"valuation/pkg/excel"
	"valuation/pkg/file"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/xuri/excelize/v2"
)

type GoodService struct {
	uc  *biz.GoodUsecase
	log *log.Helper
	pb.UnimplementedGoodServer
}

func NewGoodService(uc *biz.GoodUsecase, logger log.Logger) *GoodService {
	return &GoodService{
		uc:  uc,
		log: log.NewHelper(logger),
	}
}

func (s *GoodService) CreateGood(ctx context.Context, req *pb.CreateGoodsRequest) (*pb.CreateGoodsReply, error) {
	good, err := s.uc.CreateGood(ctx, &biz.Good{
		Name:   req.Good.Name,
		Type:   req.Good.Type,
		Unit:   req.Good.Unit,
		Price:  req.Good.Price,
		Tariff: req.Good.Tariff,
		Alias:  convertx.Chinese2Spell(req.Good.Name),
	})
	errorx.Dangerous(err)
	return &pb.CreateGoodsReply{Id: good.Id}, nil
}
func (s *GoodService) UpdateGood(ctx context.Context, req *pb.UpdateGoodsRequest) (*pb.UpdateGoodsReply, error) {
	return &pb.UpdateGoodsReply{}, nil
}
func (s *GoodService) DeleteGood(ctx context.Context, req *pb.DeleteGoodsRequest) (*pb.DeleteGoodsReply, error) {
	return &pb.DeleteGoodsReply{}, nil
}
func (s *GoodService) GetGood(ctx context.Context, req *pb.GetGoodsRequest) (*pb.GetGoodsReply, error) {
	return &pb.GetGoodsReply{}, nil
}
func (s *GoodService) ListGoods(ctx context.Context, req *pb.ListGoodsRequest) (*pb.ListGoodsReply, error) {
	pageSize, total, goods, err := s.uc.GetGoods(ctx, req.PageNum, req.PageSize, &biz.Good{})
	errorx.Dangerous(err)
	var goodsInfo []*pb.GoodInfo
	for _, good := range goods {
		goodsInfo = append(goodsInfo, &pb.GoodInfo{
			Id:     good.Id,
			Name:   good.Name,
			Type:   good.Type,
			Unit:   good.Unit,
			Price:  good.Price,
			Tariff: good.Tariff,
			Alias:  good.Alias,
		})
	}
	return &pb.ListGoodsReply{
		DataList: goodsInfo,
		PageNum:  req.PageNum,
		PageSize: pageSize,
		Total:    total,
	}, nil
}
func (s *GoodService) ListGoodsByWords(ctx context.Context, req *pb.ListGoodsByWordsRequest) (*pb.ListGoodsByWordsReply, error) {
	goods, err := s.uc.GetGoodsByWords(ctx, req.Words)
	errorx.Dangerous(err)
	var goodsInfo []*pb.GoodInfo
	for _, good := range goods {
		goodsInfo = append(goodsInfo, &pb.GoodInfo{
			Id:     good.Id,
			Name:   good.Name,
			Type:   good.Type,
			Unit:   good.Unit,
			Price:  good.Price,
			Tariff: good.Tariff,
			Alias:  good.Alias,
		})
	}
	return &pb.ListGoodsByWordsReply{Goods: goodsInfo}, nil
}

func (s *GoodService) GoodUpload(ctx http.Context) error {
	var (
		params = []map[string]string{
			{
				"id":     "0",
				"key":    "name",
				"header": "货物（劳务）名称",
				"isNull": "false",
			},
			{
				"id":     "0",
				"key":    "type",
				"header": "规格型号",
				"isNull": "false",
			},
			{
				"id":     "0",
				"key":    "unit",
				"header": "单位",
				"isNull": "false",
			},
			{
				"id":     "0",
				"key":    "price",
				"header": "金额",
				"isNull": "false",
			},
			{
				"id":     "0",
				"key":    "tariff",
				"header": "税率",
				"isNull": "false",
			},
		}
		headRow = 0
	)

	req := ctx.Request()
	f, fHeader, err := req.FormFile("file")
	errorx.Dangerous(err)
	if fHeader.Size > 1024*1024*2 {
		errorx.Bomb(500, "file size is too larger")
	} else if !strings.HasSuffix(fHeader.Filename, ".xlsx") {
		errorx.Bomb(500, "file format is error")
	}
	defer f.Close()

	//excel落盘
	filename := excel.CreateFileName()
	fExcel, err := file.Create(path.Join(".", filename))
	errorx.Dangerous(err)
	defer func() {
		_ = fExcel.Close()
		_ = file.Remove(path.Join(".", filename))
	}()
	_, err = io.Copy(fExcel, f)
	errorx.Dangerous(err)

	//获取excel表
	myExcel, err := excel.ReadMyExcel(path.Join(".", filename))
	errorx.Dangerous(err)

	//检查表头
	for _, conf := range params {
		search, err := myExcel.Search(conf["header"])
		errorx.Dangerous(err)
		if search == nil {
			errorx.Bomb(500, "tableHeader %s not found", conf["header"])
		}
		x, y, err := excelize.CellNameToCoordinates(search[0])
		errorx.Dangerous(err)
		if headRow != 0 && y != headRow {
			errorx.Bomb(500, "tableHeader format error")
		}
		headRow = y
		conf["id"] = strconv.Itoa(x)
	}

	//获取表数据
	rows, err := myExcel.Rows(headRow, params)
	errorx.Dangerous(err)

	//获取alias
	nameList := make([]string, 0, len(rows))
	for _, row := range rows {
		nameList = append(nameList, row["name"].(string))
		row["alias"] = convertx.Chinese2Spell(row["name"].(string))
	}

	//检查是否有重复name
	goods, err := data.HasGoodsByName(nameList)
	errorx.Dangerous(err)
	if len(goods) == len(rows) {
		errorx.Bomb(201, "all goods is exist")
	}

	//去除重复name
	for _, good := range goods {
		for i, row := range rows {
			if row["name"] == good.Name {
				rows[i] = rows[len(rows)-1]
				rows = rows[:len(rows)-1]
			}
		}
	}

	//插入数据库
	errorx.Dangerous(data.InsertGoodsByExcel(rows))

	return ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"data": nil,
		"msg":  "ok",
	})
}

var (
	params = []map[string]string{
		{
			"id":     "0",
			"key":    "name",
			"header": "货物（劳务）名称",
			"isNull": "false",
		},
		{
			"id":     "0",
			"key":    "type",
			"header": "规格型号",
			"isNull": "false",
		},
		{
			"id":     "0",
			"key":    "unit",
			"header": "单位",
			"isNull": "false",
		},
		{
			"id":     "0",
			"key":    "price",
			"header": "金额",
			"isNull": "false",
		},
		{
			"id":     "0",
			"key":    "tariff",
			"header": "税率",
			"isNull": "false",
		},
	}
	headRow = 0
)

func GoodUpload(ctx http.Context) error {
	req := ctx.Request()
	f, fHeader, err := req.FormFile("file")
	errorx.Dangerous(err)
	if fHeader.Size > 1024*1024*2 {
		errorx.Bomb(500, "file size is too larger")
	} else if !strings.HasSuffix(fHeader.Filename, ".xlsx") {
		errorx.Bomb(500, "file format is error")
	}
	defer f.Close()

	//excel落盘
	filename := excel.CreateFileName()
	fExcel, err := file.Create(path.Join(".", filename))
	errorx.Dangerous(err)
	defer fExcel.Close()
	_, err = io.Copy(fExcel, f)
	errorx.Dangerous(err)

	//获取excel表
	myExcel, err := excel.ReadMyExcel(path.Join(".", filename))
	errorx.Dangerous(err)

	//检查表头
	for _, conf := range params {
		search, err := myExcel.Search(conf["header"])
		errorx.Dangerous(err)
		if search == nil {
			errorx.Bomb(500, "tableHeader %s not found", conf["header"])
		}
		x, y, err := excelize.CellNameToCoordinates(search[0])
		errorx.Dangerous(err)
		if headRow != 0 && y != headRow {
			errorx.Bomb(500, "tableHeader format error")
		}
		headRow = y
		conf["id"] = strconv.Itoa(x)
	}

	//获取表数据
	rows, err := myExcel.Rows(headRow, params)
	errorx.Dangerous(err)

	//获取alias
	nameList := make([]string, 0, len(rows))
	for _, row := range rows {
		nameList = append(nameList, row["name"].(string))
		row["alias"] = convertx.Chinese2Spell(row["name"].(string))
	}

	//检查是否有重复name
	goods, err := data.HasGoodsByName(nameList)
	errorx.Dangerous(err)
	if len(goods) == len(rows) {
		errorx.Bomb(201, "all goods is exist")
	}

	//去除重复name
	for _, good := range goods {
		for i, row := range rows {
			if row["name"] == good.Name {
				rows[i] = rows[len(rows)-1]
				rows = rows[:len(rows)-1]
			}
		}
	}

	//插入数据库
	errorx.Dangerous(data.InsertGoodsByExcel(rows))

	return ctx.JSON(200, map[string]interface{}{
		"code": 200,
		"data": nil,
		"msg":  "ok",
	})
}
