def range(start, finish)
  result = []

  for i in start..finish do
    result << i
  end

  return result
end

def rangeOdd(start, finish)
  result = []

  for i in start..finish do
    if (i % 2 != 0) 
      result << i
    end
  end

  return result
  end

range(15, 30)
rangeOdd(15, 30)