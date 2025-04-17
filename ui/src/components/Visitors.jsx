import { useEffect,useState } from 'react'

const API = "http://127.0.0.1:8080";

export default function Visitors () {

    const [count,setCount] = useState();

    useEffect(()=>{
        async function getCount() {
            var response = await fetch(`${API}/visitors`);
            let data = await response.json();
            
            setCount(data.visitorCount)
        }
        getCount();
    },[]);

    return (
        <div>
            <p className='font-semibold text-lg py-4' id = "visitorCountTxt">Visitor Count : {count}</p>

        </div>
    )
}