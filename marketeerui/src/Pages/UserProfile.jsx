 import React, {UseEffect, useState} from 'react'
import axios from 'axios';

function UserProfile(){
    const [posts, setPosts] = useState([]);
    UseEffect( () => {
        axios.get('https://jsonplaceholder.typicode.com/users')
        .then(res =>{
            console.log(res)
        })
        .catch(err =>{
            console.log(err)
        })
    });


    return(
        // <div>
        //     <ul>
        //         {
        //             posts.map(post => <li key={post.id}>{post.name}</li>)
        //         }
        //     </ul>
        // </div>
        console.log('hello world')
    )
}   
export default UserProfile

// const display = () => 
// { 
//     // fetch('http://192.168.1.5/marketeer/login')  
//     fetch('http://192.168.1.5/marketeer/displayitemforsale')
//     .then((res) => res.json()) 
//     .then((data) => console.log(data))
//     .catch((err) => console.log(err))
// }
// export default display