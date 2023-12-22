const range = (start, end) => {
  let result = [];

  for (let i = start; i <= end; i += 1) {
    result.push(i);
  }

  return result;
};

const rangeOdd = (start, end) => {
  let result = [];

  for (let i = start; i <= end; i += 1) {
    if (i % 2 !== 0) {
      result.push(i);
    }
  }

  return result;
};

range(15, 30);
rangeOdd(15, 30);
