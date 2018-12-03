#!/usr/bin/env bash
#!/usr/bin/env bash


#配置build config文件路径
BuildConfigFilePkg="config"
BuildConfigFileName="buildConfig.go"
BuildConfigFile=""


#基本编译信息

Properties=()
#版本号
versionCode=1000
#版本名称
versionName="1.0.0"

#编译类型
buildType="debug"


#删除旧的文件
funDelOldFile(){
    BuildConfigFile=${BuildConfigFilePkg}"/"${BuildConfigFileName}
    if [ -f ${BuildConfigFile} ];
    then
        rm -f ${BuildConfigFile}
    fi
}

#写入包信息
funEchoPkgInfo(){
    echo "package "${BuildConfigFilePkg} >> ${BuildConfigFile}
    echo "" >> ${BuildConfigFile}
}

#写入自定义属性
funEchoProperties(){
    echo "var(" >> ${BuildConfigFile}

    for i in "${!Properties[@]}"
    do
        echo -e "\t${Properties[$i]}" >> ${BuildConfigFile}
    done

    echo ")" >> ${BuildConfigFile}
}

#添加自定义属性到数组中
funAddProperty(){
  if [ $# > 0 ]
  then
    Properties[${#Properties[*]}]=$1
  fi
}

#初始化
funInit(){
    funDelOldFile

    funAddProperty "VersionCode = ${versionCode}"
    funAddProperty "VersionName = \"${versionName}\""

    if [ $1 == "debug" ]
    then
        funAddProperty "Debug = true"
    else
        funAddProperty "Debug = false"
    fi
}


buildDebugFun(){
    go run main.go -isParseSwagger=true

    go-bindata -o=asset/asset.go -pkg=asset staticServer/...

    go run main.go
}

buildReleaseFun(){
    go run main.go -isParseSwagger=true

    go-bindata -o=asset/asset.go -pkg=asset staticServer/...

    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
}
#执行编译操作
funCreateBuildConfig(){
    funEchoPkgInfo
    funEchoProperties
}


#执行命令入口
if [ $* >0 ]
then
    buildType=$1
fi

#初始化配置信息
funInit ${buildType}

#配置自定义参数
funAddProperty "Port = \":9494\""

if [ ${buildType} == "debug" ]
then
#    funAddProperty "MySql = \"root:123456@(localhost:3306)/sense100?charset=utf8mb4&parseTime=true&loc=Local\""
    funAddProperty "MySql = \"root:xhshop123!@#@(47.106.36.243:3306)/small_program_test?charset=utf8mb4&parseTime=true&loc=Local\""
    funAddProperty "Host = \"192.168.1.104:9494\""

elif [ ${buildType} == "debug_mc" ]
then
    funAddProperty "MySql = \"root:123456@(localhost:3306)/sense100?charset=utf8mb4&parseTime=true&loc=Local\""
    funAddProperty "Host = \"192.168.1.104:9494\""

elif [ ${buildType} == "debug_ol" ]
then
    funAddProperty "MySql = \"root:xhshop123!@#@(47.106.36.243:3306)/small_program_test?charset=utf8mb4&parseTime=true&loc=Local\""
    funAddProperty "Host = \"47.106.36.243:9494\""

elif [ ${buildType} == "release" ]
then
    funAddProperty "MySql = \"root:xhshop123!@#@(47.106.36.243:3306)/small_program_test?charset=utf8mb4&parseTime=true&loc=Local\""
    funAddProperty "Host = \"47.106.36.243:9494\""
fi

#创建编译配置文件
funCreateBuildConfig

if [ ${buildType} == "debug" ]
then
    buildDebugFun
elif [ ${buildType} == "debug_mc" ]
then
    buildDebugFun
elif [ ${buildType} == "debug_ol" ]
then
    buildReleaseFun
else
    buildReleaseFun
fi


