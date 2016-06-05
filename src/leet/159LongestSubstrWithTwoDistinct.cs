using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LongestSubstrWithTwoDistinct
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine("{0}", s.LengthOfLongestSubstringTwoDistinct("eceba"));
            Console.WriteLine("{0}", s.LengthOfLongestSubstringTwoDistinct("ecebaba"));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int LengthOfLongestSubstringTwoDistinct(string s)
        {
            char? first = null;
            char? second = null;
            int lastFirst = -1;
            int lastSecond = -1;
            int curLength = 0;
            int result = 0;

            for (var i = 0; i < s.Length; i++)
            {
                if (!first.HasValue || first.Value == s[i])
                {
                    curLength++;
                    first = s[i];
                    lastFirst = i;
                }
                else if (!second.HasValue || second.Value == s[i])
                {
                    curLength++;
                    second = s[i];
                    lastSecond = i;
                }
                else
                {
                    if (lastFirst < lastSecond)
                    {
                        curLength = i - lastFirst;
                        first = s[i];
                        lastFirst = i;
                    }
                    else
                    {
                        curLength = i - lastSecond;
                        second = s[i];
                        lastSecond = i;
                    }
                }

                

                if (result < curLength)
                {
                    result = curLength;
                }
            }

            return result;
        }
    }
}
