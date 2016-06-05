using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ReverseBits
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.reverseBits(43261596));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        private bool GetBit(uint n, int ndx)
        {
            var mask = 1 << ndx;
            return (n & mask) != 0;
        }

        private uint SetBit(uint n, int ndx, bool value)
        {
            var mask = (uint)1 << ndx;
            if (value)
            {
                return n | mask;
            }

            return n & (uint.MaxValue - mask);
        }
        public uint reverseBits(uint n)
        {
            for (var i = 0; i < 16; i++)
            {
                var opNdx = 31 - i;
                var value = GetBit(n, i);
                var opValue = GetBit(n, opNdx);
                n = SetBit(n, i, opValue);
                n = SetBit(n, opNdx, value);
            }

            return n;
        }
    }
}
