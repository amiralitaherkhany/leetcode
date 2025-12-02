void main(List<String> args) {
  print(Solution().isPowerOfTwo(1024));
}

class Solution {
  bool isPowerOfTwo(int n) {
    int temp = 1;
    if (n == 1) return true;
    if (n % 2 != 0) return false;
    while (temp < n) {
      temp *= 2;
      if (temp == n) {
        return true;
      }
    }
    return false;
  }
}
