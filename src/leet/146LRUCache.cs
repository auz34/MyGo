using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LRUCache
{
    class Program
    {
        static void Main(string[] args)
        {
            var cache = new LRUCache(3);

            cache.Set(1, 1);
            cache.Set(2, 2);
            cache.Set(3, 3);
            cache.Set(4, 4);

            var get1 = cache.Get(1);
            var get2 = cache.Get(2);
            Console.WriteLine("1 should not exist anymore proof = {0}", get1);
            Console.WriteLine("2 should still exist anymore proof = {0}", get2);
            cache.Set(5, 5);
            var get3 = cache.Get(3);
            Console.WriteLine("3 should not exist anymore proof = {0}", get3);


            Console.ReadLine();
        }
    }

    public class LRUCache
    {
        private class LinkNode
        {
            public LinkNode Prev;

            public LinkNode Next;

            public int Value;

            public int key;
        }

        private int capacity;

        private LinkNode root;

        private LinkNode end;

        private Dictionary<int, LinkNode> map = new Dictionary<int, LinkNode>();

        public LRUCache(int capacity)
        {
            this.capacity = capacity;
        }

        public int Get(int key)
        {
            if (!this.map.ContainsKey(key))
            {
                return -1;
            }

            var value = this.map[key].Value;
            this.RemoveNode(key);
            this.AddAsRoot(key, value);
            return value;
        }

        public void Set(int key, int value)
        {
            if (!this.map.ContainsKey(key) && this.map.Count == this.capacity)
            {
                this.RemoveNode(this.end.key);
            }

            if (this.map.ContainsKey(key))
            {
                this.RemoveNode(key);
            }

            this.AddAsRoot(key, value);
        }

        private void RemoveNode(int key)
        {
            var node = this.map[key];
            if (node.Next != null)
            {
                node.Next.Prev = node.Prev;
            }
            else
            {
                this.end = node.Prev;
            }

            if (node.Prev != null)
            {
                node.Prev.Next = node.Next;
            }
            else
            {
                this.root = node.Next;
            }

            this.map.Remove(key);
        }

        private void AddAsRoot(int key, int value)
        {
            var node = new LinkNode { key = key, Next = root, Prev = null, Value = value};
            if (this.root != null)
            {
                this.root.Prev = node;
            }
            
            this.root = node;
            if (this.end == null)
            {
                this.end = node;
            }
            this.map.Add(key, node);
        }
    }
}
