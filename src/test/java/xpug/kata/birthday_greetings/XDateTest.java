package xpug.kata.birthday_greetings;
import static org.junit.Assert.*;

import org.junit.Test;



public class XDateTest {
	@Test
	public void getters() throws Exception {
		XDate xDate = new XDate("1789/01/24");
		assertEquals(1, xDate.getMonth());
		assertEquals(24, xDate.getDay());
	}
	
	@Test
	public void isSameDate() throws Exception {
		XDate xDate = new XDate("1789/01/24");
		XDate sameDay = new XDate("2001/01/24");
		XDate notSameDay = new XDate("1789/01/25");
		XDate notSameMonth = new XDate("1789/02/25");
		
		assertTrue("same", xDate.isSameDay(sameDay));
		assertFalse("not same day", xDate.isSameDay(notSameDay));
		assertFalse("not same month", xDate.isSameDay(notSameMonth));
	}
	
	@Test
	public void equality() throws Exception {
		XDate base = new XDate("2000/01/02");
		XDate same = new XDate("2000/01/02");
		XDate different = new XDate("2000/01/04");
		
		assertFalse(base.equals(null));
		assertFalse(base.equals(""));
		assertTrue(base.equals(base));
		assertTrue(base.equals(same));
		assertFalse(base.equals(different));
	}

}
