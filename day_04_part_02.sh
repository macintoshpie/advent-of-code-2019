#!/bin/bash

function validSequence {
  local val=$1
  local lastMatched=-1
  local matchCount=0
  local hasNeighbors=0
  for (( j=1; j < ${#val}; j++ )); do
    k=$(( $j - 1 ))
    if [[ ${val:$j:1} < ${val:$k:1} ]]; then
      return 1
    elif [[ ${val:$j:1} == ${val:$k:1} ]]; then
      # previous character matches current character
      if [[ ${val:$j:1} != $lastMatched ]]; then
        if [[ $matchCount == 2 ]]; then
          # we had a valid match
          hasNeighbors=1
        fi
        matchCount=2
        lastMatched=${val:$j:1}
      else
        matchCount=$(( $matchCount + 1 ))
      fi
    fi
  done
  if [[ $hasNeighbors == 1 || $matchCount == 2 ]]; then
    return 0
  fi
  return 1
}

function main {
  # This could probably be reduced to a combinatorics problem; might come back later
  # This could also be done more intelligently by constructing numbers with increasing digits
  read input
  echo $input
  local minVal=${input%-*}
  local maxVal=${input#*-}
  local nPasswords=0

  for (( i=$minVal; i<=$maxVal; i++ )); do
    if validSequence $i; then
      nPasswords=$(( $nPasswords + 1 ))
    fi
  done
  echo "Number of possible passwords: $nPasswords"
}

main