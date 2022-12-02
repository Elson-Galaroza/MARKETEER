
// const display = () => 
// { 
//     // fetch('http://192.168.1.5/marketeer/login')  
//     fetch('http://192.168.1.5/marketeer/systemcheck')
//     .then((res) => res.json()) 
//     .then((data) => console.log(data))
//     // .catch((err) => console.log(err))

   

// }
// export default display
import { useEffect, useState } from "react";

function ItemForSale() {
  const [data, setData] = useState([]);

  const fetchData = () => {
    fetch(`http://192.168.1.5/marketeer/displayitemforsale`)
      .then((response) => response.json())
      .then((actualData) => {
        console.log(actualData);
        setData(actualData.products);
        console.log(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  };

  useEffect(() => {
    fetchData();
  }, []);

  return (
    <div>
        <table>
            <tbody>
                <tr>
                    <th>Name</th>
                    <th>Price</th>
                    <th>Description</th>
                </tr>
                { 
                    data?.map((item) => (
                            <td key={item.ItemId}> {item.ItemName} </td>
                        )
                    )
                }
            </tbody>
        </table>
    </div>
  );
}
export default ItemForSale;