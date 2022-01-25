#!/bin/awk -f

function array_to_dec(a, len) {
  n = 0
  for (i = len; i > 0; i--) {
    n += (2 ^ (len - i)) * a[i]
  }
  return n
}

{
  len = split($0, line, "")
  for (i = 1; i <= len; i++ ) {
    if (line[i] == 1) {
      res[i] += 1
    } else {
      res[i] -= 1
    }
  }
}

END {
  for (i in res) {
    if ( res[i] > 0) {
      gamma[i] = 1
      epsilon[i] = 0
    } else {
      gamma[i] = 0
      epsilon[i] = 1
    }
  }
  gamma_dec = array_to_dec(gamma, 12)
  epsilon_dec = array_to_dec(epsilon, 12)
  printf("gamma = %d, epsilon = %d, gamma * epsilon = %d\n", gamma_dec, epsilon_dec, gamma_dec * epsilon_dec)
}
