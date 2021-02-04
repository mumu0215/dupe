'use strict';

/*
* @param  {String} pwd   连接密码
* @param  {Array}  data  编码器处理前的 payload 数组
* @return {Array}  data  编码器处理后的 payload 数组
*/
module.exports = (pwd, data, ext={}) => {
  // ##########    请在下方编写你自己的代码   ###################
  // 原有的 payload 在 data['_']中
  // 取出来之后，转为 base64 编码并放入 randomID key 下
  data[pwd] = Buffer.from(data['_']).toString('base64');
  
  //base64编码的前后再拼接随意定义的一个字符串
  data[pwd] = "Aaron" + data[pwd];
  data[pwd] += "Ruben";
  
  // ##########    请在上方编写你自己的代码   ###################

  // 删除 _ 原有的payload
  delete data['_'];
  // 返回编码器处理后的 payload 数组
  return data;
}