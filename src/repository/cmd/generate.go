package main

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"gorm.io/gen"
	"os"

	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {

	ctx := context.Background()
	outPath := os.Getenv("OUT_PATH")
	if outPath == "" {
		outPath = "src/repository/dao"
	}

	modelPkgPath := os.Getenv("MODEL_PKG_PATH")
	if modelPkgPath == "" {
		modelPkgPath = "src/model/po"
	}

	//在windows环境，是按照相对路径来识别的。
	if util.IsWindows() {
		modelPkgPath = "../model/po"
	}

	generator := gen.NewGenerator(gen.Config{
		OutPath:      outPath,
		ModelPkgPath: modelPkgPath,
		Mode:         gen.WithDefaultQuery,
	})

	klog.CtxInfof(ctx, "outPath: %v", outPath)
	klog.CtxInfof(ctx, "modelPkgPath: %v", modelPkgPath)

	db := config.DefaultMysqlConfig().Init()

	generator.UseDB(db)
	generator.ApplyBasic(
		// 表
		generator.GenerateModelAs("scm_author", "AuthorPO"),
		generator.GenerateModelAs("scm_column", "ColumnPO"),
		generator.GenerateModelAs("scm_column_quote", "ColumnQuotePO"),
		generator.GenerateModelAs("scm_commission", "CommissionPO"),
		generator.GenerateModelAs("scm_contract", "ContractPO"),
		generator.GenerateModelAs("scm_editor", "EditorPO"),
		generator.GenerateModelAs("scm_payment", "PaymentPO"),
		generator.GenerateModelAs("scm_reader", "ReaderPO"),
		generator.GenerateModelAs("scm_receipt", "ReceiptPO"),
	)

	generator.Execute()
}
