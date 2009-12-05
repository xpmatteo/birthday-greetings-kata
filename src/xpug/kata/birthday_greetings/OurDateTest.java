package xpug.kata.birthday_greetings;
import static org.junit.Assert.*;

import org.junit.Test;



public class OurDateTest {
	@Test
	public void getters() throws Exception {
		OurDate ourDate = new OurDate("1789/01/24");
		assertEquals(1, ourDate.getMonth());
		assertEquals(24, ourDate.getDay());
	}
	
	@Test
	public void isSameDate() throws Exception {
		OurDate ourDate = new OurDate("1789/01/24");
		OurDate sameDay = new OurDate("2001/01/24");
		OurDate notSameDay = new OurDate("1789/01/25");
		OurDate notSameMonth = new OurDate("1789/02/25");
		
		assertTrue("same", ourDate.isSameDay(sameDay));
		assertFalse("not same day", ourDate.isSameDay(notSameDay));
		assertFalse("not same month", ourDate.isSameDay(notSameMonth));
	}
	
	@Test
	public void equality() throws Exception {
		OurDate base = new OurDate("2000/01/02");
		OurDate same = new OurDate("2000/01/02");
		OurDate different = new OurDate("2000/01/04");
		
		assertFalse(base.equals(null));
		assertFalse(base.equals(""));
		assertTrue(base.equals(base));
		assertTrue(base.equals(same));
		assertFalse(base.equals(different));
	}

}
