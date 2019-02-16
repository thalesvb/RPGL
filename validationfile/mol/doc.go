/*
Package mol handles My Own Library catalog file.

My Own Library is a simple XML file to keep some metadata related to a archive file.
The catalog file ends with .mol.xml, simply because some editors didn't highlight
the elements when file ended with .mol.


Basic My Own Library XML file

 <?xml version="1.0" encoding="UTF-8"?>
 <MyOwnLibrary>
   <header collection="COLLECTION_NAME" />
   <archives>
     <archive file="GAME_FILE.EXT">
       <title>GAME_TITLE</title>
     </archive>
   </archives>
 </MyOwnLibrary>
*/
package mol
