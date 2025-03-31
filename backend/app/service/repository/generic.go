package repository

type Repository interface {
    List(dest interface{}, query string) error
    Get(ID string, dest interface{}, query string) error
    Create(data interface{}, query string) error
    Update(ID string, data interface{}, query string) error
}



func List(dest interface{}, query string) error {
    
    return nil
}