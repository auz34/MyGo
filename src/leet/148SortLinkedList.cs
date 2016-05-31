using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace SortLinkedList
{
    class Program
    {
        static void Main(string[] args)
        {
            var list = GetList(new List<int>() { 5, 3, 8, 1, 7, 2, 6, 4 });
            Console.WriteLine("Input:");
            PrintList(list);
            Console.WriteLine("Output:");
            var s = new Solution();
            PrintList(s.SortList(list));
            Console.ReadLine();

        }

        static ListNode GetList(IEnumerable<int> values)
        {
            ListNode root = null, last = null;

            foreach (var value in values)
            {
                var newNode = new ListNode(value);
                if (last == null)
                {
                    root = newNode;
                }
                else
                {
                    last.next = newNode;
                }

                last = newNode;
            }

            return root;
        }

        static void PrintList(ListNode node)
        {
            while (node != null)
            {
                Console.Write(" {0} ", node.val);
                node = node.next;
            }
            Console.WriteLine();
        }
    }

    public class ListNode
    {
        public int val;
        public ListNode next;
        public ListNode(int x) { val = x; }
    }

    public class Solution
    {
        private class ListInfo
        {
            public ListNode Head;

            public ListNode Tail;
        }

        private ListInfo AppendNode(ListInfo listInfo, ListNode node)
        {
            node.next = null;
            if (listInfo == null)
            {
                return new ListInfo() { Head = node, Tail = node};
            }

            listInfo.Tail.next = node;
            listInfo.Tail = node;

            return listInfo;
        }


        private ListInfo AppendList(ListNode first, ListInfo second)
        {
            if (first == null)
            {
                return second;
            }

            var root = first;
            var tail = root;
            while (tail.next != null)
            {
                tail = tail.next;
            }

            if (second != null)
            {
                tail.next = second.Head;
                tail = second.Tail;
            }

            return new ListInfo {Head = root, Tail = tail };
        }

        private ListNode AppendList(ListInfo first, ListNode second)
        {
            if (first == null)
            {
                return second;
            }

            first.Tail.next = second;

            return first.Head;
        }


        public ListNode SortList(ListNode head)
        {
            if (head == null)
            {
                return null;
            }

            ListInfo vList = null, lList = null, rList = null;
            var v = head.val;
            var curNode = head;

            while (curNode != null)
            {
                var next = curNode.next;
                if (v > curNode.val)
                {
                    lList = AppendNode(lList, curNode);
                }
                else if (v < curNode.val)
                {
                    rList = AppendNode(rList, curNode);
                }
                else
                {
                    vList = AppendNode(vList, curNode);
                }

                curNode = next;
            }

            var left = SortList(lList != null ? lList.Head : null);
            var right = SortList(rList != null ? rList.Head : null);
            var result = AppendList(left, vList);
            return AppendList(result, right);
        }
    }
}
