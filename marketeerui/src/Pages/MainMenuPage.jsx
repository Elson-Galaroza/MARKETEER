import React, {useRef, useState, useEffect} from "react";


function MainMenuPage(){
    const [itemName, setItemName] = useState(null);

const fetchData = () => {
        fetch(`http://localhost/marketeer/displayitemforsale`, {mode: "no-cors",  method: "GET"})
         .then((response)=> response.json())
         .then((data) => setItemName(data));

        console.log(test);
    }
}
export default MainMenuPage;