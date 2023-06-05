export const formatFrequency = ({ cadence, periods, startDate }) => {
  const CADENCES = {
    DAILY: "day",
    WEEKLY: "week",
    EVERY_TWO_WEEKS: "two weeks",
    MONTHLY: "month",
    EVERY_TWO_MONTHS: "two months",
    EVERY_FOUR_MONTHS: "four months",
    EVERY_SIX_MONTHS: "six months",
    ANNUAL: "year"
  };

  const newStartDate = new Date(Date.parse(startDate));

  const endDate = () => {
    let days;

    switch (cadence) {
      case "DAILY":
        newStartDate.setDate(newStartDate.getDate() + periods);
        return newStartDate.toISOString().split('T')[0];
      case "WEEKLY":
        days = periods * 7;
        newStartDate.setDate(newStartDate.getDate() + days);
        return newStartDate.toISOString().split('T')[0];
      case "EVERY_TWO_WEEKS":
        days = periods * 14;
        newStartDate.setDate(newStartDate.getDate() + days);
        return newStartDate.toISOString().split('T')[0];
      case "MONTHLY":
        newStartDate.setMonth(newStartDate.getMonth() + (1 * periods));
        return newStartDate.toISOString().split('T')[0];
      case "EVERY_TWO_MONTHS":
        newStartDate.setMonth(newStartDate.getMonth() + (2 * periods));
        return newStartDate.toISOString().split('T')[0];
      case "EVERY_FOUR_MONTHS":
        newStartDate.setMonth(newStartDate.getMonth() + (4 * periods));
        return newStartDate.toISOString().split('T')[0];
      case "EVERY_SIX_MONTHS":
        newStartDate.setMonth(newStartDate.getMonth() + (6 * periods));
        return newStartDate.toISOString().split('T')[0];
      case "ANNUAL":
        newStartDate.setFullYear(newStartDate.getFullYear() + periods);
        return newStartDate.toISOString().split('T')[0];
      default:
        break;
    }
  }

  const formattedEndDate = endDate();

  return `Every ${CADENCES[cadence]} until ${formattedEndDate}`;
};
