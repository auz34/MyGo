using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ReverseInt
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.Reverse(123));
            Console.WriteLine(s.Reverse(-123));
            Console.WriteLine(s.Reverse(-1023));
            Console.WriteLine(s.Reverse(-2147483648));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        public int Reverse(int x)
        {
            var stack = new Stack<int>();
            while (x != 0)
            {
                var n = x % 10;
                x = (x - n) / 10;
                stack.Push(n);
            }

            var result = 0;
            var mult = 1;
            while (stack.Count > 0)
            {
                try
                {
                    var part = checked(stack.Pop() * mult);
                    result = checked(result + part);
                }
                catch (OverflowException)
                {
                    return 0;
                } 
                
                mult *= 10;
            }

            return result;
        }
    }
}
