import { useState } from 'react';
import Modal from 'react-modal';
import CadencePicker from './cadencePicker';
import { useSelector } from 'react-redux';
import {
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
  const [amount, setAmount] = useState(0);
  const [description, setDescription] = useState('');
  const [payerEmail, setPayerEmail] = useState('');
  const [periods, setPeriods] = useState(0);
  const [cadence, setCadence] = useState('');
  const [startDate, setStartDate] = useState('');
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
  }

  const handleSubmit = async (event) => {
    event.preventDefault();

    const body = {
      name,
      description,
      amount,
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
            <Dialog open={modalIsOpen} onClose={closeModal}>
              <DialogTitle>Set up subscription</DialogTitle>
              <DialogContent>
                <DialogContentText>
                  Some kind of description about setting up your subscription plan.
                </DialogContentText>
                <TextField
                  id="description"
                  label="Description"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  variant="standard"
                  fullWidth
                />
                <TextField
                  id="amount"
                  label="Amount"
                  value={amount}
                  onChange={(e) => setAmount(e.target.value)}
                  variant="standard"
                  type="number"
                  inputProps={{ min: 100 }}
                  fullWidth
                />
                <TextField
                  id="payer"
                  label="From"
                  value={payerEmail}
                  onChange={(e) => setPayerEmail(e.target.value)}
                  variant="standard"
                  type="email"
                  fullWidth
                />
                <TextField
                  id="startDate"
                  label="Start Date"
                  value={startDate}
                  onChange={(e) => setStartDate(e.target.value)}
                  variant="standard"
                  type="date"
                  fullWidth
                />
                <TextField
                  id="periods"
                  label="Periods"
                  value={periods}
                  onChange={(e) => setPeriods(e.target.value)}
                  variant="standard"
                  type="number"
                  inputProps={{ min: 1 }}
                  fullWidth
                />
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
