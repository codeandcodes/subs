function CadencePicker({cadence, setCadence}) {

  const handleOptionChange = (e) => {
    setCadence(e.target.value);
  };

  return(
    <div>
      <p>Frequency</p>
      <label>
        <input
          type='radio'
          value='DAILY'
          checked={cadence === 'DAILY'}
          onChange={handleOptionChange}
        />
        Daily
      </label>
      <label>
        <input
          type='radio'
          value='WEEKLY'
          checked={cadence === 'WEEKLY'}
          onChange={handleOptionChange}
        />
        Weekly
      </label>
      <label>
        <input
          type='radio'
          value='EVERY_TWO_WEEKS'
          checked={cadence === 'EVERY_TWO_WEEKS'}
          onChange={handleOptionChange}
        />
        Every two weeks
      </label>
      <label>
        <input
          type='radio'
          value='MONTHLY'
          checked={cadence === 'MONTHLY'}
          onChange={handleOptionChange}
        />
        Monthly
      </label>
    </div>
  )
  

  // THIRTY_DAYS = 3;
  // SIXTY_DAYS = 4;
  // NINETY_DAYS = 5;
  // EVERY_TWO_MONTHS = 7;
  // QUARTERLY = 8;
  // EVERY_FOUR_MONTHS = 9;
  // EVERY_SIX_MONTHS = 10;
  // ANNUAL = 11;
  // EVERY_TWO_YEARS = 12;

}

export default CadencePicker;