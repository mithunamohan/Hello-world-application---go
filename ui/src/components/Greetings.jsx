import { useEffect,useState } from 'react'

const API = "127.0.0.1:8000";

export default function Greetings(){

    const [greetings,setGreetings] = useState(""); 
    const [ greetingsList,setGreetingsList] = useState([]);

    useEffect(()=>{
        async function getGreetings(){
            var response = await fetch(`http://${API}/greetings`);
            let data = await response.json();
            setGreetingsList(data); //store full list
            setGreetings(data[0]); // set first greeting
        }
        getGreetings();

    },[]);

    function changeGreeting(){
        const randomIndex = Math.floor(Math.random() * greetingsList.length);
        setGreetings(greetingsList[randomIndex]);
    }

    return(
        <div>
        <p className="font-bold text-2xl py-4"
        id="greetingTxt">{greetings}</p>
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-4 px-6 rounded"
        onClick= {changeGreeting}>Change Greetings</button>
        </div>
    );
}