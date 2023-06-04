import { useState } from 'react';
import Modal from 'react-modal';
import { setupSubscription } from '../../api/subscription';
import CadencePicker from './cadencePicker';
import { useSelector } from 'react-redux';

Modal.setAppElement('#root');

function SetupSubscriptionModal() {
  const [modalIsOpen, setIsOpen] = useState(false);
  const [name, setName] = useState('');
  const [amount, setAmount] = useState(0);
  const [description, setDescription] = useState('');
  const [payerEmail, setPayerEmail] = useState('');
  const [periods, setPeriods] = useState(0);
  const [cadence, setCadence] = useState('');
  const [startDate, setStartDate] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const user = useSelector(state => state.session.user);

  const tomorrow = new Date();
  tomorrow.setDate(tomorrow.getDate() + 1);

  const formattedTomorrow = tomorrow.toISOString().split('T')[0];

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
      await setupSubscription(body);
      setIsOpen(false);
    } catch (err) {
      setError(err.message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div>
      <button onClick={openModal}>Setup New Subscription</button>

      <Modal
        isOpen={modalIsOpen}
        onRequestClose={closeModal}
        contentLabel="Example Modal"
      >
        {isLoading ? (
          <div>Loading...</div>
        ) : error ? (
          <div>Error: {error}</div>
        ) : (
          // Render the form when not loading and no error
          <div>
          <button onClick={closeModal}>close</button>
          <div>New subscription</div>
          <form onSubmit={handleSubmit}>
            <div>
              <label>
                Name:
                <input
                  type="text"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </label>
            </div>
            <div>
              <label>
                description:
                <input
                  type="text"
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                />
              </label>
            </div>
            <div>
              <label>
                amount:
                <input
                  type="number"
                  value={amount}
                  min="100"
                  onChange={(e) => setAmount(e.target.value)}
                />
              </label>
            </div>
            <div>
              <p>Payer information</p>
              <div>
                  <label>
                    email address:
                    <input
                      type="string"
                      value={payerEmail}
                      onChange={(e) => setPayerEmail(e.target.value)}
                    />
                  </label>
                </div>
            </div>
            <div>
              <p>Details</p>
              <div>
                <label>
                      Start Date:
                      <input
                        type="date"
                        value={startDate}
                        onChange={(e) => setStartDate(e.target.value)}
                        // min={formattedTomorrow}
                      />
                    </label>
                </div>
              <div>
                <label>
                      periods:
                      <input
                        type="number"
                        value={periods}
                        onChange={(e) => setPeriods(e.target.value)}
                      />
                    </label>
                </div>
            </div>
            <CadencePicker cadence={cadence} setCadence={setCadence}/>
            <button type="submit">Submit</button>
          </form>
          </div>
        )}
      </Modal>
    </div>
  );
};

export default SetupSubscriptionModal;
