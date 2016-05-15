using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LongestPalyndrom
{
    using System.Globalization;

    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            var case1 = new List<int> { 1, 2, 3, 4, 5, 1 };
            var answer1 = s.FindLongestPalindrom(case1);
            Console.WriteLine("Input #1:");
            PrintArrayLine(case1);
            Console.WriteLine("Output #1:");
            PrintArrayLine(answer1);

            var case2 = new List<int> { 1, 2, 3, 4, 4, 3, 2, 1 };
            var answer2 = s.FindLongestPalindrom(case2);
            Console.WriteLine("Input #2:");
            PrintArrayLine(case2);
            Console.WriteLine("Output #2:");
            PrintArrayLine(answer2);

            var case3 = new List<int> { 1, 2, 3, 1, 4, 7, 6, 4, 3, 7, 2, 1 };
            var answer3 = s.FindLongestPalindrom(case3);
            Console.WriteLine("Input #3:");
            PrintArrayLine(case3);
            Console.WriteLine("Output #3:");
            PrintArrayLine(answer3);

            Console.ReadLine();
        }

        static void PrintArrayLine(List<int> ar)
        {
            var sb = new StringBuilder();
            for (var i=0; i<ar.Count; i++)
            {
                sb.AppendFormat(" {0} ", ar[i]);
            }

            Console.WriteLine(sb.ToString());
        }
    }

    public class Solution
    {
        private class Pair
        {
            public int First;

            public int Second;
        }

        private Pair ContinueSubsequence(
            int first,
            int firstValue,
            List<Pair> subsequence,
            Dictionary<int, List<int>> inputMap)
        {
            var pairCandidates = inputMap[firstValue];
            var lastIndexInSub = subsequence[subsequence.Count - 1].Second;
            for (var i = pairCandidates.Count - 1; i >= 0; i--)
            {
                if (pairCandidates[i] > first && pairCandidates[i] < lastIndexInSub)
                {
                    return new Pair() { First = first, Second = pairCandidates[i] };
                }
            }

            return null;
        }

        private Pair FindPair(int first, int firstValue, Dictionary<int, List<int>> inputMap)
        {
            var pairCandidates = inputMap[firstValue];
            var last = pairCandidates[pairCandidates.Count - 1];
            if (last > first)
            {
                return new Pair() {First = first, Second = last };
            }

            return null;
        }

        private List<int> GenerateAnswer(List<Pair> sequence, List<int> input)
        {
            if (sequence == null)
            {
                return new List<int>() { input[0] };
            }

            var result = new List<int>();
            for (var i = 0; i < sequence.Count; i++)
            {
                result.Add(input[sequence[i].First]);
            }

            var lastPair = sequence.Last();
            if (lastPair.First + 1 < lastPair.Second)
            {
                result.Add(input[lastPair.First + 1]);
            }

            for (var i = sequence.Count - 1; i >= 0; i--)
            {
                result.Add(input[sequence[i].Second]);
            }

            return result;
        }

        public List<int> FindLongestPalindrom(List<int> input)
        {
            var map = new Dictionary<int, List<int>>();
            for (var i = 0; i < input.Count; i++)
            {
                if (!map.ContainsKey(input[i]))
                {
                    map.Add(input[i], new List<int>());
                }

                map[input[i]].Add(i);
            }

            var subSequences = new List<List<Pair>>();

            for (var i = 0; i < input.Count; i++)
            {
                var longestIndex = -1;
                var longestSize = 0;
                Pair pairToLongest = null;
                
                for (var j = i - 1; j >= 0; j--)
                {
                    if (subSequences[j] == null)
                    {
                        continue;
                    }

                    if (subSequences[j].Count <= longestSize)
                    {
                        continue;
                    }

                    var pair = this.ContinueSubsequence(i, input[i], subSequences[j], map);
                    if (pair == null)
                    {
                        continue;
                    }

                    longestSize = subSequences[j].Count;
                    longestIndex = j;
                    pairToLongest = pair;
                }

                if (longestIndex != -1)
                {
                    var newSeq = new List<Pair>(subSequences[longestIndex]);
                    newSeq.Add(pairToLongest);
                    subSequences.Add(newSeq);
                }
                else
                {
                    var pair = FindPair(i, input[i], map);
                    if (pair != null)
                    {
                        subSequences.Add(new List<Pair> { pair });
                    }
                    else
                    {
                        subSequences.Add(null);
                    }
                }
            }

            var longestSeqSize = 0;
            List<Pair> longestSequence = null;

            foreach (var sequence in subSequences)
            {
                if (sequence == null)
                {
                    continue;
                }

                var curSize = sequence.Count * 2;
                var lastPair = sequence.Last();
                if (lastPair.First + 1 < lastPair.Second)
                {
                    curSize++;
                }

                if (curSize > longestSeqSize)
                {
                    longestSeqSize = curSize;
                    longestSequence = sequence;
                }
            }

            return this.GenerateAnswer(longestSequence, input);
        }
            
        
    }
}
