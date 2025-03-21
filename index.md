[#](#.md) Einfalda Notes
  [文档规范](zet-010221164751-73.md)
  
	
Quick note
* #vim Tips
		`:put execute('map')`  put the output of command(like 'map') into current buffer;
* [#GnuPG](220505-1448.md) 
		
#git
cherry pick
`git checkout master & git cherry-pick Commit1, Commt2 `


#Golang
# Traps

1. mix short declaration and declaration may case scope error
```go
func main(){
	var x int = 3 
	if x, y := func(){return 1, 2} {
		fmt.Printf("x: %d--y:%d", x, y) // x will print as 2
	}
	fmt.Println("x",  x) //  x will been print as 3
}
```
		
		
		
## Toc
* Language
  * [Golang](zet-310121131409-66.md)
  * [Python](zet-310121135923-71.md)
  * [Javascript/EmacScript](zet-010221201023-75.md)
  * [Java](zet-010221201057-75.md)
  * [vue](zet-310121190616-70.md)
  * [shell](zet-310121190714-70.md)
  * [CSS](210202-1732.md)
  * [Mermaid](210427-2004.md)

* Application
  * [MySQL](zet-310121140846-71.md)
  * [redis](zet-310121121417-65.md)
  * [zookeeper](zet-310121123430-65.md)
  * [docker](zet-280121172134-66.md)
  * [ElasticSearch](210203-1047.md)
  * [git](zet-310121123614-65.md)
  * [nginx](zet-310121122541-65.md)
  * [hadoop](zet-310121140733-71.md)

* Misces
  * [Authentication](210902-2319.md)
  * [Encode&Crypto](210802-1607.md)
  * [Firewall](210220-1050.md)
  * [PKI](210531-1024.md)
  * [Database](210203-2311.md)
  * [Architect](210203-1756.md)
  * [Theory](210202-1736.md)
  * [Regular Expressions](210220-1427.md)
  * [Network](210203-1417.md)
    * [Network Tools](210202-1726.md)
    * [TCP](zet-010221162407-73.md)
    * [HTTP](zet-020221161959-75.md)
    * [Network misc](210311-1720.md)
  * [Data Structure](210312-1740.md)
    * Array
    * List
    * [Tree](210203-2259.md)
    * Graph
    * [Dynamic Programing](210328-1738.md)
  * [Algorithm](210203-2257.md)
* [unspecified](210202-1737.md)
* [terminal](210630-1813.md)
* [Fangs](Fangs.md)
