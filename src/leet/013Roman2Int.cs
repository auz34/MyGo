using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Roman2Int
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine("{0}", s.RomanToInt("IIV"));
            Console.WriteLine("{0}", s.RomanToInt("XII"));
            Console.WriteLine("{0}", s.RomanToInt("XVII"));
            Console.WriteLine("{0}", s.RomanToInt("XXXVII"));
            Console.WriteLine("{0}", s.RomanToInt("IX"));
            Console.WriteLine("{0}", s.RomanToInt("DCXXI"));
            
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int RomanToInt(string s)
        {
            var sum = 0;
            var maxPrev = -1;
            for (var i = s.Length - 1; i >= 0; i--)
            {
                var cur = 0;
                switch (s[i])
                {
                    case 'I':
                        cur = 1;
                        break;
                    case 'V':
                        cur = 5;
                        break;
                    case 'X':
                        cur = 10;
                        break;
                    case 'L':
                        cur = 50;
                        break;
                    case 'C':
                        cur = 100;
                        break;
                    case 'D':
                        cur = 500;
                        break;
                    case 'M':
                        cur = 1000;
                        break;
                }

                if (cur >= maxPrev)
                {
                    sum += cur;
                }
                else
                {
                    sum -= cur;
                }

                if (cur > maxPrev)
                {
                    maxPrev = cur;
                }
            }

            return sum;
        }
    }
}
