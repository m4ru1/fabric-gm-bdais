## 更新日志

### master 更新（Sep 17, 2020）
- 私钥新增Decrypt函数，实现crypto.Decrypter接口

### master 更新（Sep 11, 2020）
- 新增导入导出接口  

   | 接口名 | 接口功能 | 
   | --- | --- |
   | PrivateKeyToPEM | 将私钥转为pem字节流 |
   | PEMtoPrivateKey 将pem字节流转成sm2私钥 | 
   | PublicKeyToPEM | 将公钥转为pem字节流 |
   | PEMtoPublicKey | 将pem字节流转为sm2公钥 |


### master 更新 （Aug 20, 2020）
- 新增性能测试数据<br>
    测试环境：
    - cpu：intel i7-7700 3.6GHz
    - 内存：16G<br>
    
   测试结果<br>
   
   | 测试算法 | 签名速度（tps） | 验签速度（tps） |
   | --- | --- | --- |
   | sm2 | 47920 | 36792 |
   | ecdsa p256 | 48832 | 16779 |

    
- 新增CHANGLOG文件


### master 更新（Aug 13, 2020）
- 优化国密tls库，支持tls双国密证书通信。

### master 更新（Aug 7, 2020）
- 将中国网安的两个底层密码库合并到一起。

### master 更新（Jul 17, 2020）
- 上传中国网安国密密码库。





