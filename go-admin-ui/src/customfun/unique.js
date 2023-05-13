// 校验
export function onlyVerification(data, keyName, rule, value, callback) {
    // console.log(rule, value, callback);
    const valid = data.map((x) => x[keyName]).find((y) => y == value);
    if (valid) {
      callback(new Error("编号已存在"));
    } else {
      callback();
    }
  }
  