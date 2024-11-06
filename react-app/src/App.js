import Select from 'react-dropdown-select';
import './App.css';
import { useState } from 'react';

function App () {
  const options = [
    {
      value: "USD",
      label: "USD"
    },
    {
      value: "CAD",
      label: "CAD"
    },
    {
      value: "MXN",
      label: "MXN"
    },
    {
      value: "EUR",
      label: "EUR"
    },
    {
      value: "GBP",
      label: "GBP"
    },
  ]

  const [currencyFrom, setCurrencyFrom] = useState("");
  const [currencyTo, setCurrencyTo] = useState("");
  const [amount, setAmount] = useState(0);

  const [convertedAmount, setConvertedAmount] = useState(undefined);

  const handleAmountChange = (event) => {
    setAmount(event.target.value);
  };

  const handleCurrencyFromChange = (values) => {
    setConvertedAmount(undefined);
    setCurrencyFrom(values[0].value);
  };

  const handleCurrencyToChange = (values) => {
    setConvertedAmount(undefined);
    setCurrencyTo(values[0].value);
  };


  const handleConversion = async () => {
    if (!currencyFrom || !currencyTo || !amount) {
      alert("need to select all values")
    }

    const conversionResponse = await fetch(`/convert?from=${currencyFrom}&to=${currencyTo}&amount=${amount}`)

    const jsonResponse = await conversionResponse.json();
    setConvertedAmount(Number(jsonResponse).toFixed(2));
  }


  const ConversionResult = () => {
    return (
      <div className='flex w-full justify-center items-center'>
        <div className='text-3xl'>{amount}</div>
        <div className='text-3xl'>&nbsp;<span className='font-extrabold'>{currencyFrom}</span>&nbsp;</div>
        <div className='text-3xl'>&rarr; <span className='font-extrabold'>{currencyTo}</span></div>
        <div className='text-3xl'> &nbsp;=&nbsp;{convertedAmount}&nbsp;<span className='font-extrabold'>{currencyTo}</span></div>
      </div>
    )
  }

  return (
    <div className="App py-14">
      <h1 className="text-3xl font-bold">
        Let's change some currency
      </h1>

      <div className='flex flex-col md:flex-row justify-center items-center w-full py-14 gap-20'>
        <div className='flex w-3/4 md:w-1/12 flex-col'>
          <label>From</label>
          <Select options={options} className='w-1/4' onChange={handleCurrencyFromChange} ></Select>
        </div>
        <div className='flex w-3/4 md:w-1/12  flex-col'>
          <label>To</label>
          <Select options={options} className='w-1/4' onChange={handleCurrencyToChange}></Select>
        </div>
        <div className='flex w-3/4 md:w-1/12  flex-col'>
          <label htmlFor='amountInput'>Amount</label>
          <input type='number' aria-label='amount' id='amountInput' className='border-solid border-2 py-1 px-2' value={amount} onChange={handleAmountChange}></input>
        </div>
      </div>

      <div className='flex justify-center w-full py-4 md:py-8'>
        <button onClick={handleConversion} className='border-2 font-extrabold uppercase text-xl border-solid rounded border-blue-500 bg-blue-400 p-4 px-12'>
          Convert &nbsp; &nbsp;&nbsp;💵
        </button>
      </div>
      {convertedAmount && <ConversionResult />}
    </div>
  );
}

export default App;
