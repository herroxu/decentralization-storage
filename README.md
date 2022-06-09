## 使用go包decentralization-storage上传文件

1. 下载decentralization-storage包

   ```go
   go get github.com/kingdemoking/decentralization-storage
   ```

2. 用户注册了账号后，可以通过登陆获取到token

   ```go
    import (
	    "fmt"
	    utils "github.com/kingdemoking/decentralization-storage/api"
    )

    func main(){
	    res := utils.GetToken("105581911@qq.com", "12356")
	    fmt.Println(res.Token)
    }
   ```
   得到的正确返回值：

   ```json
   {
     "code": 200,
     "token": "iEufSDNtJdWa.......",
     "msg": "success"
   }
   ```

3. 获取到token后，需要获取用户的存储key值

   ```go
    import (
	    "fmt"
	    utils "github.com/kingdemoking/decentralization-storage/api"
    )

    func main(){
        token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9......."
	    res := utils.GetAPIKey(token)
	    fmt.Println(res.Key)
    }
   ```

   得到的正确返回值：

   ```json
   {
     "code": 200,
     "key": "iEufSDNtJdWa.......",
     "msg": "success"
   }
   ```

4. 得到key值后，就可以进行文件存储, 目前提供了两种存储方式，Arweave/IPFS, 根据用户自己的需求，填入ipfs/arweave参数即可。

   ```go
    import (
	    "fmt"
	    utils "github.com/kingdemoking/decentralization-storage/api"
    )

    func main(){
       apiKey := "GYQ6iysjaTCo0eyvPZx3W3ggva1Ev..."
	    to := "ipfs"
	    filepath := "/home/lucas/girl1.png"

	    res := utils.Upload(apiKey, to, filepath)
	    fmt.Println(res.Url)
    }
   ```

   上传成功后，会得到返回值

   ```go
   {
     "url": "https://arweave.net/Vt6a7X0XAS0zrHKnrJvLqtDjPV9LcddKAAjVMACekg4",
     "code": 200,
     "msg": "success"
   }
   ```

6. 当上传成功后，可以通过浏览器输入url查看

   ![](https://i.postimg.cc/fTvGVXKF/iron.png)

   


