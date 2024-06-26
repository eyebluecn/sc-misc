package repo

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2po"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type PaymentRepo struct {
}

func NewPaymentRepo() PaymentRepo {
	return PaymentRepo{}
}

// 新建一个Payment
func (receiver PaymentRepo) Insert(
	ctx context.Context,
	payment *do.PaymentDO,
) (*do.PaymentDO, error) {
	table := dao.Use(config.DB).PaymentPO

	//时间置为当前
	payment.CreateTime = time.Now()
	payment.UpdateTime = time.Now()

	paymentDO := do2po.ConvertPaymentPO(payment)

	err := table.WithContext(ctx).Debug().Create(paymentDO)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}

	return po2do.ConvertPaymentDO(paymentDO), nil
}

// 更新状态
func (receiver PaymentRepo) UpdateStatus(
	ctx context.Context,
	paymentId int64,
	paymentStatus enums.PaymentStatus,
) (int64, error) {
	table := dao.Use(config.DB).PaymentPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(paymentId))

	info, err := table.WithContext(ctx).Debug().Where(conditions...).Update(table.Status, paymentStatus)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return 0, err
	}

	return info.RowsAffected, nil
}

// 按照id查找，找不到返回nil
func (receiver PaymentRepo) QueryById(
	ctx context.Context,
	paymentId int64,
) (*do.PaymentDO, error) {
	table := dao.Use(config.DB).PaymentPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(paymentId))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	paymentDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}
	return po2do.ConvertPaymentDO(paymentDO), nil
}

// 按照id查找，找不到返回NotFound的Err
func (receiver PaymentRepo) CheckById(
	ctx context.Context,
	paymentId int64,
) (*do.PaymentDO, error) {
	payment, err := receiver.QueryById(ctx, paymentId)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, errs.CodeErrorf(errs.StatusCodeUnknown, err.Error())
	}
	if payment == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "没有找到%v对应的记录", paymentId)
	}
	return payment, nil
}

// 按照id查找，找不到返回nil
func (receiver PaymentRepo) QueryByOrderNo(
	ctx context.Context,
	orderNo string,
) (*do.PaymentDO, error) {
	table := dao.Use(config.DB).PaymentPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.OrderNo.Eq(orderNo))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	paymentDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}
	return po2do.ConvertPaymentDO(paymentDO), nil
}

// 按照id查找，找不到返回NotFound的Err
func (receiver PaymentRepo) CheckByOrderNo(
	ctx context.Context,
	orderNo string,
) (*do.PaymentDO, error) {
	payment, err := receiver.QueryByOrderNo(ctx, orderNo)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, errs.CodeErrorf(errs.StatusCodeUnknown, err.Error())
	}
	if payment == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "没有找到orderNo=%v对应的记录", orderNo)
	}
	return payment, nil
}
