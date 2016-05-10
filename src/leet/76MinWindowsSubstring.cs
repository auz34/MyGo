using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace MinWindowSubstring
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = "ADOBECODEBANC";
            var t = "ABC";
            var solution = new Solution();
            Console.WriteLine("Use Case #1: {0}", solution.MinWindow(s, t));
            Console.WriteLine("Use Case #2: {0}", solution.MinWindow("a", "aa"));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        private string GetWindow(string s, List<List<int>> candidates, List<int> pointers, out int minCandidateNdx)
        {
            var min = s.Length;
            var max = -1;
            minCandidateNdx = -1;
            for (var i = 0; i < candidates.Count; i++)
            {
                var ndx = candidates[i][pointers[i]];
                if (ndx < min)
                {
                    min = ndx;
                    minCandidateNdx = i;
                }

                if (ndx > max)
                {
                    max = ndx;
                }
            }

            return s.Substring(min, max - min + 1);
        }

        public string MinWindow(string s, string t)
        {
            if (s.Length < t.Length)
            {
                return string.Empty;
            }

            var map = new Dictionary<char, List<int>>();
            foreach (var ch in t)
            {
                if (!map.ContainsKey(ch))
                {
                    map.Add(ch, new List<int>());
                }
            }

            for (var i = 0; i < s.Length; i++)
            {
                if (map.ContainsKey(s[i]))
                {
                    map[s[i]].Add(i);
                }
            }

            var candidates = map.Values.ToList();
            if (candidates.Any(l => l.Count == 0))
            {
                return string.Empty;
            }

            var pointers = new List<int>();
            for (var i = 0; i < candidates.Count; i++)
            {
                pointers.Add(0);
            }

            int minCandidateNdx;
            var result = this.GetWindow(s, candidates, pointers, out minCandidateNdx);
            while (pointers[minCandidateNdx] + 1 < candidates[minCandidateNdx].Count)
            {
                pointers[minCandidateNdx]++;
                var resCandidate = this.GetWindow(s, candidates, pointers, out minCandidateNdx);
                if (resCandidate.Length < result.Length)
                {
                    result = resCandidate;
                }
            }

            return result;
        }
    }
}
