using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LongestSubstring
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine("{0}", s.LengthOfLongestSubstring("abcabcbb"));
            Console.WriteLine("{0}", s.LengthOfLongestSubstring("bbbbb"));
            Console.WriteLine("{0}", s.LengthOfLongestSubstring("pwwkew"));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int LengthOfLongestSubstring(string s)
        {
            var map = new Dictionary<char, List<int>>();
            var ndxMap = new List<int>();
            for (var i = 0; i < s.Length; i++)
            {
                if (!map.ContainsKey(s[i]))
                {
                    map.Add(s[i], new List<int>());
                }

                var charMap = map[s[i]];
                charMap.Add(i);
                ndxMap.Add(charMap.Count - 1);
            }

            var prev = 0;
            var result = 0;
            for (var i = 0; i < s.Length; i++)
            {
                var indexInMap = ndxMap[i];
                if (indexInMap == 0)
                {
                    prev++;
                }
                else
                {
                    var charMap = map[s[i]];
                    var prevOccurence = charMap[indexInMap - 1];
                    prev = Math.Min(i - prevOccurence - 1, prev) + 1;
                }

                if (prev > result)
                {
                    result = prev;
                }
            }

            return result;
        }
    }
}
