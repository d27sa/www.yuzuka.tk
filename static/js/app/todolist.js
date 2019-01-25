function newRequest(){
    return new XMLHttpRequest();
}

window.onload=function(){
    var deletes=document.getElementsByClassName('delete');
    for(var i=0;i<deletes.length;i++){
        deletes[i].onclick=function(){
            var li =this.parentElement;
            var req = newRequest();
            req.open('POST',document.location.origin+'/app/todolist/'+this.getAttribute('href'));
            req.onload=function(){
                if(req.status==200){
                    li.parentElement.removeChild(li);
                }
            };
            var todo=li.getElementsByTagName('span')[0];
            req.send(todo.innerText);
            return false;
        }
    }
};