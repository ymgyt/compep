======
compep
======


compep is a my utility package for competitive programming in Go.


usage
=====

lines
-----

.. code-block:: go

   lines := NewLines(os.Stdin, 4096)
   n := lines.Next().Int()

   for i := 0; i < n; i++ {
    ss := lines.Next().Strings()
    // handle input
   }