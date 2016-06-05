using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace NumberOf1Bits
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.HammingWeight(5));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int HammingWeight(uint n)
        {
            var result = 0;
            while (n > 0)
            {
                var k = n % 2;
                n = n >> 1;
                if (k == 1)
                {
                    result++;
                }
            }

            return result;
        }
    }
}
