package main
import(
	"net/http"
	"text/template"
	"log"
	"os"
)
var tpl *template.Template
type Todo struct {
    Name string
	Recommanded_For string
	Maintenance_level string
	Lifespan string
	Temperament string
	Health_risk string
	Image_url string
    Url  string
	Img_one string
	Img_two string
	Img_three string
}



type TodoPageData struct {
    Todos     []Todo
}
func init(){
	tpl=template.Must(template.ParseFiles("form.html","animal.html","information.html"))
}
func main(){
	port:=os.Getenv("PORT")
	
	http.HandleFunc("/",indexfunc)
	http.HandleFunc("/animal",animal_list)
	http.HandleFunc("/info",animal_info)
	log.Fatal(http.ListenAndServe(":"+port,nil))
	
}
func indexfunc(w http.ResponseWriter,r *http.Request){
	tpl.Execute(w,nil)
}
func animal_list(w http.ResponseWriter,r *http.Request){

	data := TodoPageData{
            
            Todos: []Todo{
				{Name: "Airedale Terrier", Recommanded_For : "Small Families", Maintenance_level : "High", Lifespan : "10-14 years", Temperament: "Intelligent", Health_risk : "High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure.",Image_url : "https://images2.minutemediacdn.com/image/upload/c_fill,g_auto,h_1248,w_2220/f_auto,q_auto,w_1100/v1555312680/shape/mentalfloss/istock_79171327_small.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/airedale-terrier/",Img_one: "data:image/jpeg;base64,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",Img_two:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSvD4vsQZlYu4iS4OYP0-UlGILBKb43KN11UA&usqp=CAU",Img_three:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTQywd8qTO40DPOX-zZWExjAO9zZrJiibHeZQ&usqp=CAU",},
				
				{Name: "Akita", Recommanded_For :"Small families, singles (wary of strangers)", Maintenance_level : "High", Lifespan : "11-15 years", Temperament: "Active", Health_risk :"High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure." ,Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/10/akita-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/akita/",Img_one:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTh0mTOAOFwnT6FU30YNmTDmlk1Ab5SaTK5bQ&usqp=CAU",Img_two:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQs38K3OUL1_M63Kwm0j51yR048x_d8qtYSdA&usqp=CAU",Img_three :"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTuEH3n8KmgTO7e0xwxP4rJYWFg-nh4qzZkJQ&usqp=CAU",},
				//{Name: "Task 3", Recommanded_For :, Maintenance_level : , Lifespan : , Temperament: , Health_risk : ,Image_url : , Url: "true"},
				{Name: "Alaskan Malamute", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "12-16 years", Temperament: "Intelligent", Health_risk : "High probability of health issues during its lifetime, hence it is one of the more expensive breeds to insure.",Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/10/alaskan-malamute-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/alaskan-malamute/",Img_one:"data:image/jpeg;base64,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",Img_two:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRe0LINigTm6RZINeUNC5O47V5IJaKrAK667g&usqp=CAU",Img_three:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQkwsXjyEguNgmkIwPX8jmFNimVfY_pWT8QhA&usqp=CAU",},
				{Name: "American Bulldog", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "8-15 years", Temperament: "friendly", Health_risk :"Higher than average probability of developing health issues during its lifetime, hence the cost to insure is above average." ,Image_url : "https://bowwowinsurance.com.au/wp-content/uploads/2018/11/american-bulldog-700x700.jpg", Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/american-bulldog/",Img_one:"data:image/jpeg;base64,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",Img_two:"https://media.gettyimages.com/photos/american-bulldog-picture-id516156469?s=612x612",Img_three:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcRpqjghs0Uw4lorgyqcXdxX4NtXJTdZZDnZ6w&usqp=CAU",},
				{Name: "American Staffordshire Terrier", Recommanded_For :"Families", Maintenance_level : "Medium", Lifespan : "10-12 years", Temperament: "social", Health_risk : "This breed has an around average probability of having health issues in its lifetime, hence it is one of the more affordable breeds to insure.",Image_url :"https://bowwowinsurance.com.au/wp-content/uploads/2018/10/american-staffordshire-terrier-amstaff-american-staffy-700x700.jpg" , Url: "https://bowwowinsurance.com.au/dogs/dog-breeds/american-staffordshire-terrier/",Img_one:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcSGSdm3JhXNcuWbvHkOYpwekTyjrmiOz681jA&usqp=CAU",Img_two :"https://www.dogster.com/wp-content/uploads/2016/11/shutterstock_104234978-600x792.jpg",Img_three:"https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQm1fpAIaYYvNae2ipKZcmMbp5BCtLXKKMfOg&usqp=CAU",},
            },
        }
	tpl.ExecuteTemplate(w,"animal.html",data)
}
func animal_info(w http.ResponseWriter,r *http.Request){
	tpl.ExecuteTemplate(w,"information.html",nil)
}