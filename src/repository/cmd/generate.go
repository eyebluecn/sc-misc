package main

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"gorm.io/gen"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {

	ctx := context.Background()
	outPath := os.Getenv("OUT_PATH")
	if outPath == "" {
		outPath = "src/repository/query"
	}

	modelPkgPath := os.Getenv("MODEL_PKG_PATH")
	if modelPkgPath == "" {
		modelPkgPath = "db_model"
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:      outPath,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery,
	})

	klog.CtxInfof(ctx, "outPath: %v", outPath)
	klog.CtxInfof(ctx, "modelPkgPath: %v", modelPkgPath)

	db := config.InitMySQL()

	generator.UseDB(db)
	generator.ApplyBasic(
		// 表
		generator.GenerateModelAs("scm_author", "AuthorDO"),
		generator.GenerateModelAs("scm_column", "ColumnDO"),
		generator.GenerateModelAs("scm_column_quote", "ColumnQuoteDO"),
		generator.GenerateModelAs("scm_commission", "CommissionDO"),
		generator.GenerateModelAs("scm_contract", "ContractDO"),
		generator.GenerateModelAs("scm_editor", "EditorDO"),
		generator.GenerateModelAs("scm_payment", "PaymentDO"),
		generator.GenerateModelAs("scm_reader", "ReaderDO"),
		generator.GenerateModelAs("scm_receipt", "ReceiptDO"),
	)

	generator.Execute()
}
