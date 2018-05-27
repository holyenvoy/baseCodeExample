#include <iostream>
 using namespace std;
 class A
 {
 public:
    virtual ~A() { cout << "Destroy A" << endl; }
 };
 
 class B : public A
 {
 public:
   ~B() { cout << "Destroy B" << endl; }
 };
 
 int main()
 {
   A* p = new B;
   delete p;
   return 0;
}