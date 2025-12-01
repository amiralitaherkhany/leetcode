void main(List<String> args) {
  print(Solution().isHappy(2));
}

class Solution {
  int times = 0;
  bool isHappy(int n) {
    int result = 0;
    String number = n.toString();
    for (var i = 0; i < number.length; i++) {
      result += (int.parse(number[i]) * int.parse(number[i]));
    }
    if (result == 1) {
      return true;
    }
    if (times > 10) {
      return false;
    }
    times++;
    return isHappy(result);
  }
}
