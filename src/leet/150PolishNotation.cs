using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace PolishNotation
{
    class Program
    {
        static void Main(string[] args)
        {
            var s = new Solution();

            var case1 = new string[] { "2", "1", "+", "3", "*" };
            var case2 = new string[] { "4", "13", "5", "/", "+" };

            Console.WriteLine(s.EvalRPN(case1));
            Console.WriteLine(s.EvalRPN(case2));

            Console.ReadLine();
        }
    }

    public class Solution
    {
        private bool isOperation(string token)
        {
            return token == "+" || token == "-" || token == "*" || token == "/";
        }

        private int GetOperand(string token)
        {
            return int.Parse(token);
        }

        private int ApplyOp(int operand1, int operand2, string operation)
        {
            switch (operation)
            {
                case "+":
                    return operand1 + operand2;
                case "-":
                    return operand1 - operand2;
                case "*":
                    return operand1 * operand2;
                case "/":
                    return operand1 / operand2;
            }

            throw new Exception("Incorrect operation");
        }

        public int EvalRPN(string[] tokens)
        {
            var stack = new Stack<int>();

            foreach (var token in tokens)
            {
                if (this.isOperation(token))
                {
                    var op2 = stack.Pop();
                    var op1 = stack.Pop();
                    stack.Push(this.ApplyOp(op1, op2, token));
                }
                else
                {
                    stack.Push(this.GetOperand(token));
                }
            }

            return stack.Pop();
        }
    }
}
