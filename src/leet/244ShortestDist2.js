/**
 * @constructor
 * @param {string[]} words
 */
var WordDistance = function(words) {
    this.map = {};

    for (var i=0; i<words.length; i++) {
        if (!this.map.hasOwnProperty(words[i])) {
            this.map[words[i]] = [i];
        } else {
            this.map[words[i]].push(i);
        }
    }
};

WordDistance.prototype.dist = function(x, y) {
    if (x>y) {
        return x- y;
    }

    return y - x;
};

WordDistance.prototype.findMinDistance = function(ind1, ind2) {
    var d = this.dist(ind1[0], ind2[0]);

    for (var i=0; i<ind1.length; i++) {
        for (var j=0; j<ind2.length; j++) {
            if (this.dist(ind1[i], ind2[j]) < d) {
                d = this.dist(ind1[i], ind2[j]);
            }
        }
    }

    return d;
};

/**
 * @param {string} word1
 * @param {string} word2
 * @return {integer}
 */
WordDistance.prototype.shortest = function(word1, word2) {
    return this.findMinDistance(this.map[word1], this.map[word2]);
};

/**
 * Your WordDistance object will be instantiated and called as such:
 * var wordDistance = new WordDistance(words);
 * wordDistance.shortest("word1", "word2");
 * wordDistance.shortest("anotherWord1", "anotherWord2");
 */