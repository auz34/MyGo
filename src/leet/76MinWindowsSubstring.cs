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
            Console.WriteLine("Use Case #3: {0}", solution.MinWindow("ab", "aa"));
            Console.WriteLine("Use Case #4: {0}", solution.MinWindow("aa", "aa"));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        private class Pointers
        {
            private string input;

            private List<List<int>> candidates;

            private List<int> requiredCount;

            private List<int> curPointers = new List<int>();

            private int minPointerNdx;

            private int minInputNdx;
            private int maxInputNdx;

            public Pointers(string input, List<List<int>> candidates, List<int> requiredCount)
            {
                this.input = input;
                this.candidates = candidates;
                this.requiredCount = requiredCount;
                this.Initialize();
            }

            public bool IsValid { get; private set; }

            public string CurrentResult
            {
                get
                {
                    if (!this.IsValid)
                    {
                        return string.Empty;
                    }

                    return this.input.Substring(this.minInputNdx, this.maxInputNdx - this.minInputNdx + 1);
                }
            }

            public void Next()
            {
                this.curPointers[this.minPointerNdx]++;
                this.FindBorders();
            }

            private void FindBorders()
            {
                this.minInputNdx = this.input.Length;
                this.maxInputNdx = -1;

                this.IsValid = true;

                for (var i = 0; i < this.candidates.Count; i++)
                {
                    var iReq = this.requiredCount[i];
                    var iCandidates = this.candidates[i];
                    var iPointer = this.curPointers[i];
                    if (iCandidates.Count < iPointer + iReq)
                    {
                        this.IsValid = false;
                        return;
                    }

                    this.curPointers.Add(0);
                    for (var j = iPointer; j < iPointer + iReq; j++)
                    {
                        if (iCandidates[j] < this.minInputNdx)
                        {
                            this.minInputNdx = iCandidates[j];
                            this.minPointerNdx = i;
                        }

                        if (iCandidates[j] > this.maxInputNdx)
                        {
                            this.maxInputNdx = iCandidates[j];
                        }
                    }
                }
            }

            private void Initialize()
            {
                for (var i = 0; i < this.candidates.Count; i++)
                {
                    this.curPointers.Add(0);
                }

                this.FindBorders();
            }
        }

        public string MinWindow(string s, string t)
        {
            if (s.Length < t.Length)
            {
                return string.Empty;
            }

            #region Iterate through t to account its characters and count
            var charMap = new Dictionary<char, List<int>>();
            var countMap = new Dictionary<char, int>();
            foreach (var ch in t)
            {
                if (!charMap.ContainsKey(ch))
                {
                    charMap.Add(ch, new List<int>());
                }

                if (!countMap.ContainsKey(ch))
                {
                    countMap.Add(ch, 0);
                }

                countMap[ch]++;
            }
            #endregion

            #region iterate through input string to collect where characters from t are
            for (var i = 0; i < s.Length; i++)
            {
                if (charMap.ContainsKey(s[i]))
                {
                    charMap[s[i]].Add(i);
                }
            }
            #endregion

            var candidates = new List<List<int>>();
            var requiredCount = new List<int>();
            foreach (var kvp in charMap)
            {
                candidates.Add(kvp.Value);
                requiredCount.Add(countMap[kvp.Key]);
            }

            var pointers = new Pointers(s, candidates, requiredCount);
            var result = "";
            while (pointers.IsValid)
            {
                if (string.IsNullOrEmpty(result) || result.Length > pointers.CurrentResult.Length)
                {
                    result = pointers.CurrentResult;
                }
                pointers.Next();
            }

            return result;
        }
    }
}
