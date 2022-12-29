[![Go Report Card](https://goreportcard.com/badge/github.com/dmoose/slidingMinMax)](https://goreportcard.com/report/github.com/dmoose/slidingMinMax)

# slidingMinMax

A generic struct to calculate a windowed min and max of any comparable values.
The algoritm maintains a sorted deque of items plus the index at which they were entered.
The index is used to enforce the window size.  This operation is O(1) since each item is
only pushed or popped once and the max/min is always the first element.

# Algorithm
#### below describes calculating max, invert comparisons to calculate min
- Create deque of size K
- As items arrive there are two constraints to check
- (1) If there are items in queue and the first item's index is earlier than window size we remove it
- (2) If new item is larger than the item at back of queue remove all elements from back of queue that are smaller than new value.
- After constraints are met push new value onto end of queue
- The maximum value will always be at the front of the queue


## Getting Started

see tests for examples of how to create struct for specific type and add elements to it.  The min and max are returned after each call to push.
### Prerequisites

none

### Installing



## Running the tests

go test


## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/dmoose/slidingMinMax/tags).

## Authors

* **Jeff Shumate** - *Initial work* 

See also the list of [contributors](https://github.com/dmoose/slidingMinMax/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE) file for details

## Acknowledgments

* Hat tip to anyone whose code was used

