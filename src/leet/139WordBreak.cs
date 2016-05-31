using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace WordBreak
{
    using System.Runtime.CompilerServices;

    class Program
    {
        static void Main(string[] args)
        {
            var dict = new HashSet<string>() { "leet", "code" };
            var s = new Solution();
            Console.WriteLine(s.WordBreak("leetcode", dict));
            Console.WriteLine(s.WordBreak("leetccode", dict));

            var secondDict = new HashSet<string>()
                                 {
                                     "a",
                                     "aa",
                                     "aaa",
                                     "aaaa",
                                     "aaaaa",
                                     "aaaaaa",
                                     "aaaaaaa",
                                     "aaaaaaaa",
                                     "aaaaaaaaa",
                                     "aaaaaaaaaa"
                                 };

            var str = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab";


            Console.WriteLine(s.WordBreak(str, secondDict));



            Console.ReadLine();
        }
    }

    public class Solution
    {
        public bool WordBreak(string s, ISet<string> wordDict)
        {
            if (string.IsNullOrEmpty(s) || wordDict.Contains(s))
            {
                return true;
            }

            var mem = new List<bool>();
            for (var i = 0; i < s.Length; i++)
            {
                mem.Add(wordDict.Contains(s.Substring(0, i+1)));
            }

            for (var i = 1; i < s.Length; i++)
            {
                if (!mem[i - 1])
                {
                    continue;
                }

                for (var j = i; j < s.Length; j++)
                {
                    if (mem[j])
                    {
                        continue;
                    }

                    mem[j] = wordDict.Contains(s.Substring(i, j - i + 1));
                }
            }

            return mem[s.Length - 1];
        }
    }
}
