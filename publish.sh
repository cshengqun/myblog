#!/bin/sh

PROJECT=myblog
rm -rf release/${PROJECT}
mkdir -p release/${PROJECT}/log
cp -r bin release/${PROJECT}
cp -r tools release/${PROJECT}
cp -r conf release/${PROJECT}
cp -r templates release/${PROJECT}
cp -r static release/${PROJECT}


cp release/${PROJECT}/tools/restart.sh_pro release/${PROJECT}/tools/restart.sh
cp release/${PROJECT}/conf/myblog.json_pro release/${PROJECT}/conf/myblog.json
if [[ $1 = "first" ]]
then
	scp -r release/${PROJECT} ubuntu@111.231.232.47:/home/ubuntu/app/
else
    scp ./bin/myblog ubuntu@111.231.232.47:/home/ubuntu/app/myblog/bin
fi

