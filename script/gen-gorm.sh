# current script dir.
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
echo "SCRIPT_DIR: ${SCRIPT_DIR}"
# project home dir.
HOME_DIR=$(dirname ${SCRIPT_DIR})
echo "HOME_DIR: ${HOME_DIR}"

# 设置生成目标文件地址
export OUT_PATH=${HOME_DIR}/src/repository/query
# export MODEL_PKG_PATH=${HOME_DIR}/src/repository/model

# 自动生成repository中的model和query文件夹
go run ${HOME_DIR}/src/repository/cmd/generate.go
