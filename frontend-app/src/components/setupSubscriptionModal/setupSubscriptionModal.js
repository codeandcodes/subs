import { useState } from 'react';
import Modal from 'react-modal';
import CadencePicker from './cadencePicker';
import { useSelector } from 'react-redux';
import {
  Box,
  Button,
  Dialog,
  DialogTitle,
  DialogContent,
  DialogContentText,
  DialogActions,
  TextField
} from '@mui/material';
import { useDispatch } from 'react-redux';
import { addNewSubscription } from '../../store/subscription';

Modal.setAppElement('#root');

function SetupSubscriptionModal() {
  const dispatch = useDispatch();
  const [modalIsOpen, setIsOpen] = useState(false);
  const [name, setName] = useState('defaultName');
  const [amount, setAmount] = useState(undefined);
  const [description, setDescription] = useState('');
  const [payerEmail, setPayerEmail] = useState('');
  const [periods, setPeriods] = useState(undefined);
  const [cadence, setCadence] = useState('');
  const [startDate, setStartDate] = useState(undefined);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const user = useSelector(state => state.session.user);

  const today = new Date();
  const formattedToday = today.toISOString().split('T')[0];

  const openModal = () => {
    setIsOpen(true);
  }

  const closeModal = () => {
    setIsOpen(false);
    resetFields();
  }

  const resetFields = () => {
    setAmount(undefined);
    setDescription('');
    setPayerEmail('');
    setPeriods(undefined);
    setCadence('');
    setStartDate(undefined)
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    if (amount && periods && startDate && description.length > 0 &&
        payerEmail.length > 0 && cadence > 0) {
          const body = {
            name,
            description,
            amount: amount * 100,
            frequency: {
              cadence,
              startDate,
              periods
            },
            payee: [
              {
                emailAddress: user.emailAddress
              }
      
            ],
            payer: [
              {
                emailAddress: payerEmail
              }
            ]
          };
      
          setIsLoading(true);
          setError(null);
          try {
            dispatch(addNewSubscription(body));
            setIsOpen(false);
          } catch (err) {
            setError(err.message);
          } finally {
            setIsLoading(false);
            resetFields();
          }
      }
  };

  return (
    <div>
      {isLoading ? (
          <div>Loading...</div>
        ) : error ? (
          <div>Error: {error}</div>
        ) : (
          <div>
            <Button variant="outlined" size="small" onClick={openModal}>Setup New Subscription</Button>
            <Dialog open={modalIsOpen} onClose={closeModal} maxWidth="sm">
              <DialogTitle fontWeight="600">Set up a new subscription</DialogTitle>
              <DialogContent>
                <DialogContentText paddingBottom="8px">Request recurring payment from:</DialogContentText>
                <TextField
                  id="payer"
                  label="Email (of payer)"
                  value={payerEmail}
                  onChange={(e) => setPayerEmail(e.target.value)}
                  type="email"
                  size="small"
                  fullWidth
                  sx={{ paddingBottom: "8px "}}
                />
                <TextField
                  id="description"
                  label="Description"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  variant="outlined"
                  size="small"
                  sx={{ paddingBottom: "8px" }}
                  fullWidth
                />
                <TextField
                  id="amount"
                  label="Amount"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  variant="outlined"
                  type="number"
                  inputProps={{ min: 1.00, step: 0.01, inputMode: 'numeric', pattern: '[0-9]*' }}
                  placeholder="1.00"
                  size="small"
                  sx={{ paddingBottom: "8px" }}
                  fullWidth
                />
                <Box display="flex">
                  <TextField
                    id="startDate"
                    label="Start Date"
                    value={startDate}
                    onChange={(e) => setStartDate(e.target.value)}
                    variant="outlined"
                    type="date"
                    size="small"
                    defaultValue={formattedToday}
                    InputLabelProps={{ shrink: true }}
                    fullWidth
                  />
                  <TextField
                    id="periods"
                    label="Periods"
                    value={periods}
                    onChange={(e) => setPeriods(e.target.value)}
                    variant="outlined"
                    type="number"
                    size="small"
                    inputProps={{ min: 1 }}
                    sx={{ paddingLeft: "12px" }}
                    fullWidth
                  />
                </Box>
                <CadencePicker cadence={cadence} setCadence={setCadence}/>
                <DialogActions>
                  <Button onClick={closeModal}>Cancel</Button>
                  <Button onClick={handleSubmit}>Submit</Button>
                </DialogActions>
              </DialogContent>
            </Dialog>
          </div>
        )}
    </div>
  );
};

export default SetupSubscriptionModal;
