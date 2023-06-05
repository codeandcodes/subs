import {
  FormControl,
  FormControlLabel,
  FormLabel,
  Radio,
  RadioGroup
} from '@mui/material';

function CadencePicker({cadence, setCadence}) {

  const handleOptionChange = (e) => {
    setCadence(e.target.value);
  };

  return(
    <div>
      <FormControl sx={{ paddingTop: "8px" }}>
        <FormLabel id="cadence-group">Frequency</FormLabel>
        <RadioGroup
          aria-labelledby="cadence-group"
          name="cadence"
          value={cadence}
          sx={{
            '& .MuiSvgIcon-root': {
              fontSize: 16,
            },
            padding: "8px"
          }}
          onChange={handleOptionChange}
        >
          <FormControlLabel value="DAILY" control={<Radio />} label="Daily" />
          <FormControlLabel value="WEEKLY" control={<Radio />} label="Weekly" />
          <FormControlLabel value="EVERY_TWO_WEEKS" control={<Radio />} label="Every two weeks" />
          <FormControlLabel value="MONTHLY" control={<Radio />} label="Monthly" />
          <FormControlLabel value="EVERY_TWO_MONTHS" control={<Radio />} label="Every two months" />
          <FormControlLabel value="EVERY_FOUR_MONTHS" control={<Radio />} label="Every four months" />
          <FormControlLabel value="EVERY_SIX_MONTHS" control={<Radio />} label="Every six months" />
          <FormControlLabel value="ANNUAL" control={<Radio />} label="Every year"         sx={{
          '& .MuiSvgIcon-root': {
            fontSize: 16,
          },
        }} />
        </RadioGroup>
      </FormControl>
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