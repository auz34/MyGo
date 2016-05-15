using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LongestIncreasing
{
    class Program
    {
        static void Main(string[] args)
        {
            var case1 = new int[] { 10, 9, 2, 5, 3, 7, 101, 18 };
            var case2 = new int[] { 2, 2 };
            var s = new Solution();
            Console.WriteLine("{0}", s.LengthOfLIS(case1));
            Console.WriteLine("{0}", s.LengthOfLIS(case2));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int LengthOfLIS(int[] nums)
        {
            if (nums.Length == 0)
            {
                return 0;
            }

            var mem = new List<List<int>>();
            mem.Add(new List<int>() { nums[0] });
            var answer = 1;

            for (var i = 1; i < nums.Length; i++)
            {
                var longestPrevNdx = -1;
                var longestLength = 0;
                #region search for longest we can append
                for (var j = i - 1; j >= 0; j--)
                {
                    if (nums[i] <= nums[j])
                    {
                        continue;
                    }

                    if (mem[j].Count > longestLength)
                    {
                        longestLength = mem[j].Count;
                        longestPrevNdx = j;
                    }
                }
                #endregion

                if (longestPrevNdx == -1)
                {
                    mem.Add(new List<int>() { nums[i] });
                }
                else
                {
                    var newSeq = new List<int>(mem[longestPrevNdx]);
                    newSeq.Add(nums[i]);
                    mem.Add(newSeq);
                    if (newSeq.Count > answer)
                    {
                        answer = newSeq.Count;
                    }
                }
            }

            return answer;

        }
    }
}
