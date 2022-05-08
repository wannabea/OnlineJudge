PATH='/home/admin/go/src/github.com/wannabea/OnlineJudge'

cd $PATH/user && go run . &
cd $PATH/announcement && go run . &
cd $PATH/oj_api && go run .
