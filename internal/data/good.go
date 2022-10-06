package data

import (
	"context"
	"strings"

	"valuation/internal/biz"
	"valuation/pkg/storage"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type goodRepo struct {
	data *Data
	log  *log.Helper
}

func NewGoodRepo(data *Data, logger log.Logger) biz.GoodRepo {
	return &goodRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type Good struct {
	Id     uint64  ` json:"id,omitempty"`
	Name   string  ` json:"name,omitempty"`
	Type   string  ` json:"type,omitempty"`
	Unit   string  ` json:"unit,omitempty"`
	Price  float32 ` json:"price,omitempty"`
	Tariff float32 ` json:"tariff,omitempty"`
	Alias  string  ` json:"alias,omitempty"`
}

// GetGoodsByWords select * from goods where name like '%words%' or alias like '%words%'
func (g *goodRepo) GetGoodsByWords(ctx context.Context, words string) (goodList []*biz.Good, err error) {
	var goods []Good
	res := g.data.db.Where("name LIKE ? or alias LIKE ?", "%"+words+"%", "%"+words+"%").Find(&goods)
	if res.Error != nil {
		return nil, res.Error
	}

	for _, good := range goods {
		goodList = append(goodList, &biz.Good{
			Id:     good.Id,
			Name:   good.Name,
			Type:   good.Type,
			Unit:   good.Unit,
			Price:  good.Price,
			Tariff: good.Tariff,
			Alias:  good.Alias,
		})
	}
	return
}

func (g *goodRepo) GetGoods(ctx context.Context, pageNum, pageSize uint64, goodInfo *biz.Good) (total uint64, goodList []*biz.Good, err error) {
	var (
		goods []Good
		db    = g.data.db.Table("goods")
		t     int64
	)

	if strings.Trim(goodInfo.Name, " ") != "" {
		db.Where("name LIKE ? ", "%"+goodInfo.Name+"%")
	}

	if res := db.Count(&t); res.Error != nil {
		return 0, nil, res.Error
	}
	// select * from goods where id >= (select id from goods limit 20, 1) limit 20
	if res := db.Where("id >= (?)", g.data.db.Table("goods").Select("id").Limit(1).Offset(int((pageNum-1)*pageSize))).Limit(int(pageSize)).Find(&goods); res.Error != nil {
		return 0, nil, res.Error
	}

	for _, good := range goods {
		goodList = append(goodList, &biz.Good{
			Id:     good.Id,
			Name:   good.Name,
			Type:   good.Type,
			Unit:   good.Unit,
			Price:  good.Price,
			Tariff: good.Tariff,
			Alias:  good.Alias,
		})
	}

	total = uint64(t)
	return
}

func (g *goodRepo) Delete(ctx context.Context, good *biz.Good) (*biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodRepo) Save(ctx context.Context, goodInfo *biz.Good) (*biz.Good, error) {
	good := Good{
		Id:     goodInfo.Id,
		Name:   goodInfo.Name,
		Type:   goodInfo.Type,
		Unit:   goodInfo.Unit,
		Price:  goodInfo.Price,
		Tariff: goodInfo.Tariff,
		Alias:  goodInfo.Alias,
	}

	res := g.data.db.Create(&good)
	if res.Error != nil {
		return nil, res.Error
	}
	return &biz.Good{
		Id:     good.Id,
		Name:   good.Name,
		Type:   good.Type,
		Unit:   good.Unit,
		Price:  good.Price,
		Tariff: good.Tariff,
		Alias:  good.Alias,
	}, nil
}

func (g *goodRepo) Update(ctx context.Context, good *biz.Good) (*biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodRepo) FindByID(ctx context.Context, u uint64) (*biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g *goodRepo) ExistByName(ctx context.Context, s string) (bool, error) {
	var cnt int64
	if err := g.data.db.Table("goods").Where("name = ?", s).Count(&cnt).Error; err != nil {
		return false, err
	}
	if cnt != 0 {
		return true, nil
	}
	return false, nil
}

func (g *goodRepo) ListAll(ctx context.Context) ([]*biz.Good, error) {
	//TODO implement me
	panic("implement me")
}

func InsertAndUpdateGoodsByExcel(insertDate, updateData []map[string]interface{}) (RowsAffected int64, err error) {
	err = storage.DB.Transaction(func(tx *gorm.DB) error {
		if insertDate != nil {
			insertRes := tx.Table("goods").Model(&Good{}).Create(insertDate)
			if insertRes.Error != nil {
				return insertRes.Error
			}
			RowsAffected = insertRes.RowsAffected
		}

		for _, good := range updateData {
			updateRes := tx.Table("goods").Model(&Good{}).Where("name = ?", good["name"]).Update("price", good["price"])
			if updateRes.Error != nil {
				return updateRes.Error
			}
			RowsAffected++
		}
		return nil
	})

	if err != nil {
		return 0, err
	}

	return
}

func ExistsNameList(nameList []string) (exists []string, err error) {
	res := storage.DB.Table("goods").Model(&Good{}).Select("name").Where(map[string]interface{}{"name": nameList}).Find(&exists)
	if res.Error != nil {
		return
	}
	return
}

func FindPriceByName(name string) (price float64, err error) {
	res := storage.DB.Table("goods").Model(&Good{}).Select("price").Where(map[string]interface{}{"name": name}).Find(&price)
	if res.Error != nil {
		return
	}
	if res.RowsAffected == 0 {
		return -1, nil
	}
	return
}
