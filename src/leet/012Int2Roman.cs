using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Int2Roman
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();
            Console.WriteLine(s.IntToRoman(1990));
            Console.WriteLine(s.IntToRoman(1954));
            Console.WriteLine(s.IntToRoman(621));
            Console.WriteLine(s.IntToRoman(12));
            Console.WriteLine(s.IntToRoman(17));
            Console.WriteLine(s.IntToRoman(37));
            Console.WriteLine(s.IntToRoman(9));
            Console.WriteLine(s.IntToRoman(3));
            Console.ReadLine();
        }
    }

    public class Solution
    {
        private int addLetter(StringBuilder sb, int num, string letter, int letterValue)
        {
            while (num >= letterValue)
            {
                sb.Append(letter);
                num -= letterValue;
            }

            return num;
        }

        public string IntToRoman(int num)
        {
            var sb = new StringBuilder();

            num = this.addLetter(sb, num, "M", 1000);
            //num = this.addLetter(sb, num, "IM", 999);
            //num = this.addLetter(sb, num, "VM", 995);
            //num = this.addLetter(sb, num, "XM", 990);
            //num = this.addLetter(sb, num, "LM", 950);
            num = this.addLetter(sb, num, "CM", 900);

            num = this.addLetter(sb, num, "D", 500);
            //num = this.addLetter(sb, num, "ID", 499);
            //num = this.addLetter(sb, num, "VD", 495);
            //num = this.addLetter(sb, num, "XD", 490);
            //num = this.addLetter(sb, num, "LD", 450);
            num = this.addLetter(sb, num, "CD", 400);


            num = this.addLetter(sb, num, "C", 100);
            //num = this.addLetter(sb, num, "IC", 99);
            //num = this.addLetter(sb, num, "VC", 95);
            num = this.addLetter(sb, num, "XC", 90);


            num = this.addLetter(sb, num, "L", 50);
            //num = this.addLetter(sb, num, "IL", 49);
            //num = this.addLetter(sb, num, "VL", 45);
            num = this.addLetter(sb, num, "XL", 40);

            num = this.addLetter(sb, num, "X", 10);
            num = this.addLetter(sb, num, "IX", 9);

            num = this.addLetter(sb, num, "V", 5);
            num = this.addLetter(sb, num, "IV", 4);

            this.addLetter(sb, num, "I", 1);

            return sb.ToString();
        }
    }
}
