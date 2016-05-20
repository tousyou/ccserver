package matchserver
import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
)
type PeerMysql struct{
    Host    string
    Port    int32
    User    string
    Passwd  string
    conn    *sql.DB
}

func NewPeerMysql(host string,port int32,user string,passwd string) (*PeerMysql,error){
    str := fmt.Sprintf("%s:%s@tcp(%s:%d)/mytest?autocommit=true&charset=utf8",
                        user,passwd,host,port)
    fmt.Println(str)
    //db,err := sql.Open("mysql","root:12321@tcp(localhost:3306)/mytest?autocommit=true&charset=utf8")
    db,err := sql.Open("mysql",str)
    if err != nil {
        return nil,err
    }
    return &PeerMysql{
        Host:      host,
        Port:      port,
        User:      user,
        Passwd:    passwd,
        conn:      db,
    },nil
}

func (db *PeerMysql)Add(p *PeerInfo)bool{
    stmt,err := db.conn.Prepare("insert into redis(host,port) values(?,?)")
    if err != nil {
        return false
    }
    res,err := stmt.Exec(p.Host,p.Port)
    if err != nil {
        return false
    }
    affect, err := res.RowsAffected()
    if err != nil {
        return false
    }
    if affect != 1 {
        return false
    }
    return true
}
func (db *PeerMysql)Del(p *PeerInfo)bool{
    return true
}
func (db *PeerMysql)GetAll()Peers{
    return nil
}
