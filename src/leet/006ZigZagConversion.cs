using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ZigZagConversion
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.Convert("PAYPALISHIRING", 3));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public string Convert(string s, int numRows)
        {
            if (numRows == 1)
            {
                return s;
            }

            var result = new List<string>();
            for (var i = 0; i < numRows; i++)
            {
                result.Add(string.Empty);
            }

            var j = 0;
            var curRow = 0;
            var dir = 1;
            while (j < s.Length)
            {
                result[curRow] += s[j];
                curRow += dir;
                if (curRow == numRows)
                {
                    curRow -= 2;
                    dir = -1;
                } else if (curRow == -1)
                {
                    curRow = 1;
                    dir = 1;
                }

                j++;
            }

            return string.Join(string.Empty, result);
        }
    }
}
